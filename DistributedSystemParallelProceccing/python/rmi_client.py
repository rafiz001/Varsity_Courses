import Pyro5.api       # Core Pyro5 API for creating client-side Proxy stubs
import Pyro5.errors    # Pyro5 exception classes (equivalent to RemoteException)


try:
    try:
        ns = Pyro5.api.locate_ns()                  # Connect to the Name Server
        uri = ns.lookup("HelloService")             # Look up "HelloService" by name
        print(f"Found service via Name Server: {uri}")
    except Exception:
        print(f"Name Server not found.")

    # Create a Proxy — the client-side stub object.
    # Calling methods on it transparently forwards them to the remote server.
    with Pyro5.api.Proxy(uri) as hello_service:

        # Call the remote method.
        result = hello_service.say_hello()

        print(f"Server says: {result}")

except Exception as e:
    print(f"Error: {e}")

