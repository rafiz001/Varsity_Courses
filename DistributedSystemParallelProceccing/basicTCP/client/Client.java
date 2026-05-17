import java.net.*;
import java.io.*;
 
public class Client {
    public static void main(String[] args) throws IOException {
        Socket socket = new Socket("localhost", 5000);
        System.out.println("Connected to server at " + "localhost" + ":" + 5000);
        socket.close();
        System.out.println("Disconnected.");
    }
}