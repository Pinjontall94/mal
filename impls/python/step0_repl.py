import sys


def READ(x):
    return x


def EVAL(x):
    return x


def PRINT(x):
    return x


def rep(x):
    res_read = READ(x)
    res_eval = EVAL(res_read)
    res_print = PRINT(res_eval)
    return res_print


def main():
    while True:
        try:
            user_input = input("user> ")
            res_rep = rep(user_input)
            print(res_rep)
        except EOFError:
            sys.exit("\nbye! :3")


if __name__ == "__main__":
    main()
