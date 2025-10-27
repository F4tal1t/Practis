from cryptography.fernet import Fernet
# Generates a random encryption key 
def gen_key():
    return Fernet.generate_key()

# Saves the generated key into a file
def save_key(key,key_file):
    with open(key_file,'wb') as file:
        file.write(key)

# Load the encryption key to a file
def load_key(key_file):
    with open(key_file,'rb') as file:
        return file.read()

# Encrypt the file
def encrypt_file(input_file,output_file,key):
    with open(input_file,'rb') as file:
        data = file.read()
    fernet = Fernet(key)
    encrypted_data= fernet.encrypt(data)

    with open(output_file,'wb') as file:
        file.write(encrypted_data)

# Decrypt the file
def decrypt_file(input_file,output_file,key):
    with open(input_file,'rb') as file:
        encrypted_data = file.read()
    fernet = Fernet(key)
    decrypted_data= fernet.decrypt(encrypted_data)

    with open(output_file,'wb') as file:
        file.write(decrypted_data)

if __name__ == "__main__":
    key = gen_key()
    key_file='Encryptedmoment.key'
    save_key(key,key_file)

    input_file = 'getencryptedlmao.txt'
    encrypted_file='GotEncrypted.txt'
    decrypted_file='GotDecrypted.txt'

encrypt_file(input_file,encrypted_file,key)
print(f"File{input_file} is encrypted to {encrypted_file}")
decrypt_file(encrypted_file,decrypted_file,key)
print(f"File{encrypted_file} is decrypted to {decrypted_file}")
