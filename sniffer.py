import sys, socket, random
from builtins import set

arguments = sys.argv
if "--help" in arguments:
    print("Enter the host IP address and the port range. You may use domain name instead of IP address and you don't have to enter the port range(it will be set by default).")
    exit()

if "--host" not in arguments:
    print("Enter the IP address of the host!!!")
    exit()

host_ip = arguments[arguments.index("--host") + 1]
ports = """0-65535"""
if "--ports" in arguments:
    ports = arguments[arguments.index("--ports") + 1]

ports = list(map(int, ports.split("-")))
read_list = []

open_ports = []
# print(server_socket.gettimeout())
for i in range(ports[0], ports[1] + 1):
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.settimeout(0.3)
    res = server_socket.connect_ex((host_ip, i))
    if res == 0:
        print(".", end="")
        open_ports.append(str(i))
    else:
        continue
    server_socket.close()

print("\n" + ",".join(open_ports) + " ports are opened")