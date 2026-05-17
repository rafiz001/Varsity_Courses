#include <bits/stdc++.h>
using namespace std;
unordered_map<char, char> getRandomTable() {
    unordered_map<char, char> subTable;
    unordered_map<char, bool> allowedLetters;
    for (char i = 'n'; i <= 'z'; i++) {
        allowedLetters[i] = true;
    }
    for (char i = 'a'; i <= 'm'; i++) {
        while (true) {
            long nowInMilli =
            chrono::duration_cast<chrono::milliseconds>(chrono::system_clock::now().time_since_epoch()).count();
            char val = 'n' + nowInMilli%13;
            if (allowedLetters[val] == true) {
                subTable[i] = val;
                allowedLetters[val] = false;
                break;
            }
        }
    }
    return subTable;
}
string macRandom(unordered_map<char, char>& subTable, string text) {
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
    unordered_map<char, char> subTable = getRandomTable();
    string encrypted = macRandom(subTable, input);
    string decrypted = macRandom(subTable, encrypted);
    cout << "Encrypted: " << encrypted << endl;
    cout << "Decrypted: " << decrypted << endl;
    return 0;
}