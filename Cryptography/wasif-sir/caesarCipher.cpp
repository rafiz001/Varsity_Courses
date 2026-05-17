#include<bits/stdc++.h>
using namespace std;

string encrypt(const string& text, int key) {
    string result = "";
    for (char ch : text) {
        if (isalpha(ch)) {
            bool is_upper = isupper(ch);
            char keyed_char = ((ch - (is_upper ? 'A' : 'a') + key) % 26 + 26) % 26 + (is_upper ? 'A' : 'a');
            result += keyed_char;
        } else {
            result += ch;
        }
    }
    return result;
}

string decrypt(const string& ciphertext, int key) {
    return encrypt(ciphertext, -key);
}

int main() {
    string plaintext;
    cout<<"Write a message: ";
    getline(cin, plaintext);
    cout<<"Enter a secret key: ";
    int key_value;
    cin>>key_value;
    string encrypted_text = encrypt(plaintext, key_value);
    string decrypted_text = decrypt(encrypted_text, key_value);

    cout << "Original Text: " << plaintext << endl;
    cout << "Encrypted Text: " << encrypted_text << endl;
    cout << "Decrypted Text: " << decrypted_text << endl;

    return 0;
}