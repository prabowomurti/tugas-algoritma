#include <iostream>

using namespace std;

void soal01() {
    float X ;
    int A, B, C;
    A = 25; 
    B = 12;
    C = 5;
    X = A / B + C % 2 ;

    cout << "Soal 01" << endl;
    cout << "X = " << X << endl;
}

void soal02() {
    float X;
    int A, B, C;
    A = 25; 
    B = 12;
    C = 5;
    X = A / B + C / 2.0 ;

    cout << "Soal 02" << endl;
    cout << "X = " << X << endl;
}

void soal03() {
    float X ;
    int A, B, C;
    A = 25; 
    B = 12;
    C = 5;
    X = A / ( B + C) % 2 ;

    cout << "Soal 03" << endl;
    cout << "X = " << X << endl;
}

void soal04() {
    float X ;
    int A, B, C;
    A = 25; 
    B = 12;
    C = 5;
    X = A * (B + C) + 2.0 ;

    cout << "Soal 04" << endl;
    cout << "X = " << X << endl;

}

int main()
{
    soal01();
    soal02();
    soal03();
    soal04();

    return EXIT_SUCCESS;
}