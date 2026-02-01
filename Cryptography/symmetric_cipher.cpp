#include <bits/stdc++.h>
#include <cctype>
using namespace std;

string encrypt(const string& text, int key) {
    string result = "";

    for (char ch : text) {
        if (isalpha(ch)) {
            // Determine the case (uppercase or lowercase)
            bool is_upper = isupper(ch);

            // Key the character (Caesar Cipher logic)
            char keyed_char =
                ((ch - (is_upper ? 'A' : 'a') + key + 26) % 26)
                + (is_upper ? 'A' : 'a');

            result += keyed_char;
        } else {
            // Non-alphabetic characters remain unchanged
            result += ch;
        }
    }

    return result;
}

string decrypt(const string& ciphertext, int key) {
    // Decryption is encryption with a negative key
    return encrypt(ciphertext, -key);
}

int main() {
    string plaintext;
    int key_value;

    cout << "Write a message: ";
    getline(cin, plaintext);

    cout << "Enter a secret key: ";
    cin >> key_value;

    string encrypted_text = encrypt(plaintext, key_value);
    string decrypted_text = decrypt(encrypted_text, key_value);

    cout << "\nOriginal Text  : " << plaintext << endl;
    cout << "Encrypted Text : " << encrypted_text << endl;
    cout << "Decrypted Text : " << decrypted_text << endl;

    return 0;
}