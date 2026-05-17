import java.net.*;
import java.io.*;
 
public class Server {
    public static void main(String[] args) throws IOException {
        int port = 5000;
        ServerSocket ss = new ServerSocket(port);
        System.out.println("Server started. Waiting for client on port " + port + "...");
        Socket s = ss.accept(); 
        System.out.println("Client connected: " + s.getInetAddress());
        s.close();
        ss.close();
        System.out.println("Connection closed.");
    }
}