#include <bits/stdc++.h>
using namespace std;
string lowToUpper(string input) {
  string output;
  for (int i = 0; input[i]; i++) {
    if (input[i] > 'z' || input[i] < 'a') {
        output += input[i]; continue;
    }
    output += char(input[i] - ('a' - 'A'));
  }
  return output;
}
int main() {
  string input;
  cin >> input;
  cout <<lowToUpper(input)<< endl;
}