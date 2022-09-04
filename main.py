import json


def main():
    str = """
cat << EOF > pull-req-merge-commit.txt
a

b

c
EOF
  """.strip()
    print(json.dumps(str))


if __name__ == "__main__":
    main()
