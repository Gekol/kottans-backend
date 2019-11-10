import sys, os

def main():
    invalid_data = False
    try:
        file = open("counter.txt")
        num = file.read()
        if num.isdigit():
            num = int(num) + 1
            file = open("counter.txt", "w")
            file.write(str(num))
            file.close()
        else:
            invalid_data = True
    except:
        if not invalid_data:
            file = open("counter.txt", "w")
            file.write("1")
            file.close()
            os.chmod("counter.txt", 0o777)
    if invalid_data:
        sys.exit(1)

if __name__=="__main__":
    main()