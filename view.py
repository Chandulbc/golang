import sys

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python your_script.py <input>")
        sys.exit(1)

    input_value = sys.argv[1]
    print(f"Received input: {input_value}")
