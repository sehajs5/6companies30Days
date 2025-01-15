#include<iostream>
#include<string>
using namespace std;
struct TrieNode {
    TrieNode* child[26];
    bool isTerminal;

    TrieNode() {
        isTerminal = false;
        for (int i = 0; i < 26; i++) {
            child[i] = nullptr;
        }
    }
};

class Trie {
private:
    TrieNode* root;

public:
    Trie() {
        root = new TrieNode();
    }

    void insertKey(const std::string& key) {
        TrieNode* curr = root;
        for (char c : key) {
            int index = c - 'a';
            if (curr->child[index] == nullptr) {
                curr->child[index] = new TrieNode();
            }
            curr = curr->child[index];
        }
        curr->isTerminal = true;
    }

    void printSuggestions(TrieNode* curr, vector <std::string >& temp, std::string prefix) {
        if (curr->isTerminal) {
            temp.push_back(prefix);
        }
        for (char ch = 'a'; ch <= 'z'; ch++) {
            int index = ch - 'a';
            if (curr->child[index] != nullptr) {
                prefix.push_back(ch);
                printSuggestions(curr->child[index], temp, prefix);
                prefix.pop_back();
            }
        }
    }

    vector<vector<string>> getSuggestions(const string& s) {
        TrieNode* prev = root;
        vector<vector<string>> result;
        string prefix = "";

        for (char ch : s) {
            prefix.push_back(ch);
            int index = ch - 'a';

            if (prev->child[index] == nullptr) {
                break;
            }

            vector<string> temp;
            printSuggestions(prev->child[index], temp, prefix);
            result.push_back(temp);
            prev = prev->child[index];
        }

        // Fill empty suggestions for remaining characters if no match
        while (result.size() < s.size()) {
            result.push_back({"0"});
        }

        return result;
    }
};

class Solution{
public:
    vector<vector<string> > displayContacts(int n, string contact[], string s)
    {
        // code here
        Trie* t = new Trie();
        
        //insert all contact in trie
        
        for (int i = 0; i< n; i++){
            string c = contact[i];
            t->insertKey(c);
        }
        return t->getSuggestions(s);
    }
};
