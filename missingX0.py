"""
math problems version
"""

import random


def create_problem():
    data = [i + i * i for i in range(1, 10)]
    print(
        f"In the following ordered list, you will find characters or numbers ranging from {data[0]} to {data[-1]}.\n"
    )

    # Generate missingX
    x = random.choice(data)

    # Show a problem
    for v in data:
        if v != x:
            print(v, end=" ")
    ans = int(input("\n\nWhat character or number is missing? \nYour answer: "))

    return x, ans


def check_answer(x, ans):
    print("Correct!\n") if ans == x else print("Try again!")


x, ans = create_problem()
check_answer(x, ans)
