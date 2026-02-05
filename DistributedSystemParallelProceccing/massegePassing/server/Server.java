import java.net.*;
import java.io.*;

public class Server {

    public static void main(String args[]) {

        try {
            ServerSocket ss = new ServerSocket(5000);
            System.out.println("Server started");

            System.out.println("Waiting for a client ...");

            Socket s = ss.accept();
            System.out.println("Client accepted");
            DataInputStream in = new DataInputStream(s.getInputStream());
            DataOutputStream out = new DataOutputStream(s.getOutputStream());
            BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
            String m = ""; String m2 = "";

             while (!m.equals("Over")) {

                 m = in.readUTF();
                 System.out.println(m);

                m2 = br.readLine();
                out.writeUTF(m2);
                out.flush();


            }
            in.close();
            out.close();
            s.close();
            ss.close();
        } catch (Exception e) {
            System.out.println(e);
        }

    }
}