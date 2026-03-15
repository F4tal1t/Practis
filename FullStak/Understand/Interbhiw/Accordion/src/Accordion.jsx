import { useState } from "react";

import './styles.css'

function Accordion({ items }) {
    const [openIndex,setOpenindex] = useState(null)
    const toggleAccord = (index) => {
        setOpenindex(openIndex==index ? null : index)
    }
    return !items || items.length===0 ?"No items available":(
        <div className= "accordion">
            {items.map((item, index) => {
                return <div className="accordion-item">
                    <button className="accordion-title"
                        onClick={() => toggleAccord(index)}>{item.title}
                    </button>
                    {openIndex === index && <div className="accordion-content">
                        {item.content}
                    </div>}
                </div>
            })}
        </div>
    );
}

export default Accordion;