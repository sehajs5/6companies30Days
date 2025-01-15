#include<iostream>
using namespace std;
struct Node {
    int data;
    Node *next;

    Node(int x) {
        data = x;
        next = NULL;
    }
};
class Solution {
  public:
    Node* linkdelete(Node* head, int n, int m) {
        // code here
        Node* temp = head;
        while (temp!=NULL){
            for (int i=1; i<m && temp!=NULL; i++){
                temp = temp->next;
            }
            if (temp==NULL){
                break;
            }
            Node* buff = temp->next;
            for (int i=0; i<n && buff!=NULL; i++){
                Node* next = buff->next;
                delete buff;
                buff=next;
            }
            temp->next = buff;
            
            temp = buff;
        }
        
        return head;
    }
};