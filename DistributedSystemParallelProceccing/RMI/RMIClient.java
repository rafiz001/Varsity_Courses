import java.rmi.registry.*;       


public class RMIClient {
    public static void main(String [] args) {
        try {
            Registry registry = LocateRegistry.getRegistry("localhost", 5099);
            RemoteInterface helloservice = (RemoteInterface) registry.lookup("HelloService");
            String string = helloservice.sayHello();
            System.out.println("Server says: " + string);
        } catch(Exception e) {
            System.out.println(e);
        }
    }
}

