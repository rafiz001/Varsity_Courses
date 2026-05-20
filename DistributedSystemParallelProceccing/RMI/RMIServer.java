import java.rmi.*;                         
import java.rmi.server.*;  
import java.rmi.registry.*;  

public class RMIServer extends UnicastRemoteObject implements RemoteInterface {
    public RMIServer() throws RemoteException {};
    @Override
    public String sayHello() throws RemoteException {
        return "Hello World!";
    }

    public static void main(String [] args) {
        try {
            RMIServer rmiserver = new RMIServer();
            Registry registry = LocateRegistry.createRegistry(5099);
            registry.bind("HelloService", rmiserver);
            System.out.println("Server is waiting for remote call...");
        } catch(Exception e) {
            System.out.println(e);
        }
    }
}
