#include <bits/stdc++.h>
using namespace std;
string mac(string text) {
    unordered_map<char, char> subTable = {
        {'a', 'q'},
        {'b', 'w'},
        {'c', 'x'},
        {'d', 'r'},
        {'e', 't'},
        {'f', 'y'},
        {'g', 'u'},
        {'h', 's'},
        {'i', 'o'},
        {'j', 'p'},
        {'k', 'n'},
        {'l', 'v'},
        {'m', 'z'}
    };
    string result = "";
    for (char letter : text) {
        auto found = subTable.find(letter);
        if (found != subTable.end()) {
            result += found->second;
        } else {      
            for (auto& pair : subTable) {
                if (pair.second == letter) {
                    result += pair.first;
                    break;
                }
            }
        }
    }
    return result;
}
int main() {
    string input;
    cout << "Enter Text to be encrypted:" << endl;
    cin >> input;
    string encrypted = mac(input);
    string decrypted = mac(encrypted);
    cout << "Encrypted: " << encrypted << endl;
    cout << "Decrypted: " << decrypted << endl;
    return 0;
}