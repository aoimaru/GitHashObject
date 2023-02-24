# coding: utf-8
# Your code here!

def main():
    SN = list(input())
    SN.reverse()

    Ans = 0
    for SI, S in enumerate(SN):
        Num = ord(S)-64
        Ans += Num*(26**SI)

    print(Ans)

if __name__ == "__main__":
    main()