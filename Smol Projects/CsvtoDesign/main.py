import csv
from reportlab.lib.pagesizes import letter, landscape
from reportlab.pdfgen import canvas
from reportlab.lib import colors
from reportlab.platypus import Paragraph
from reportlab.lib.styles import getSampleStyleSheet, ParagraphStyle
from reportlab.lib.enums import TA_LEFT


def create_compact_pdf_from_csv(csv_file, pdf_file):
    # Read CSV data
    people = []
    try:
        # Try multiple encodings
        for encoding in ['utf-8', 'utf-8-sig', 'latin-1']:
            try:
                with open(csv_file, 'r', encoding=encoding) as file:
                    reader = csv.reader(file)
                    try:
                        next(reader)  # Skip header if present
                    except StopIteration:
                        continue

                    for row in reader:
                        if len(row) >= 6:  # Ensure we have enough columns
                            name = row[3].strip()
                            university = row[5].strip()
                            if name and university:  # Only add if both fields have data
                                people.append((name, university))
                break  # If we get here, reading succeeded
            except UnicodeDecodeError:
                continue
    except Exception as e:
        print(f"Error reading CSV: {e}")
        return

    if not people:
        print("No valid data found in the CSV file.")
        return

    # Sort people alphabetically by name
    people.sort(key=lambda x: x[0].lower())

    # Create PDF in landscape orientation for more horizontal space
    c = canvas.Canvas(pdf_file, pagesize=landscape(letter))
    width, height = landscape(letter)

    # Define box dimensions
    box_width = 180
    box_height = 60
    margin_x = 20
    margin_y = 20
    columns = 4
    rows_per_page = 9

    # Create paragraph styles for text wrapping with left alignment and Helvetica font
    styles = getSampleStyleSheet()
    name_style = ParagraphStyle(
        'NameStyle',
        parent=styles['Normal'],
        fontName='Helvetica-Bold',
        fontSize=10,
        leading=12,
        alignment=TA_LEFT
    )

    univ_style = ParagraphStyle(
        'UnivStyle',
        parent=styles['Normal'],
        fontName='Helvetica',
        fontSize=9,
        leading=11,
        alignment=TA_LEFT
    )

    # Calculate spacing - zero padding between boxes
    h_spacing = 0  # Zero horizontal spacing
    v_spacing = 0  # Zero vertical spacing

    # Adjust column positions with zero padding
    box_width = (width - 2 * margin_x) / columns

    # Create boxes with no gaps
    items_per_page = columns * rows_per_page
    total_pages = (len(people) + items_per_page - 1) // items_per_page

    for page in range(total_pages):
        start_index = page * items_per_page
        end_index = min(start_index + items_per_page, len(people))

        for i in range(start_index, end_index):
            # Calculate position
            idx = i - start_index
            col = idx % columns
            row = idx // columns

            x = margin_x + col * box_width
            y = height - margin_y - (row * box_height) - box_height

            # Draw box
            c.setStrokeColor(colors.black)
            c.setLineWidth(0.5)
            c.setFillColor(colors.white)
            c.rect(x, y, box_width, box_height, fill=1)

            # Get data
            name, university = people[i]

            # Create Paragraph objects for text wrapping
            name_para = Paragraph(name, name_style)
            univ_para = Paragraph(university, univ_style)

            # Calculate available width with padding
            left_padding = 10  # 10px padding from the left edge
            avail_width = box_width - (left_padding + 5)  # 5px padding on right side

            # Save canvas state before drawing paragraphs
            c.saveState()

            # Draw name with text wrapping
            name_width, name_height = name_para.wrap(avail_width, box_height + 60 / 2)
            name_para.drawOn(c, x + left_padding, y + box_height - 15 - name_height)

            # Draw university with text wrapping
            univ_width, univ_height = univ_para.wrap(avail_width, box_height + 30 / 2)
            univ_para.drawOn(c, x + left_padding, y + box_height / 2 - univ_height)

            # Restore canvas state
            c.restoreState()

        if page < total_pages - 1:
            c.showPage()

    c.save()
    print(f"PDF successfully created: {pdf_file} with {total_pages} pages containing {len(people)} sorted names")


# Usage
create_compact_pdf_from_csv('Her.csv', 'Participants.pdf')