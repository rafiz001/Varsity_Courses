import Pyro5.api          # Core Pyro5 API: Daemon, expose, serve
import Pyro5.nameserver   # For locating/starting the Name Server (like RMI Registry)
from remote_interface import RemoteInterface


@Pyro5.api.expose
class RMIServer(RemoteInterface):
    def say_hello(self) -> str:
        return "Hello World!"

with Pyro5.api.Daemon(host="localhost",port=5099) as daemon:
    try:
        ns = Pyro5.api.locate_ns()
        uri = daemon.register(RMIServer, "HelloService")
        ns.register("HelloService", uri)
        print(f"Registered with Name Server as 'HelloService'")
    except Exception:
        print(f"Name Server not found.")
        

    # Confirm the server is ready — equivalent to:
    # System.out.println("Server is waiting for remote call...");
    print("Server is waiting for remote call...")
    print(f"Listening on port 5099  |  Object URI: {uri}")

    daemon.requestLoop()


