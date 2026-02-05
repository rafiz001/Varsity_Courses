import java.io.*;
import java.net.*;

public class Client {

    public static void main(String[] args) {


        try {
            Socket s = new Socket("127.0.0.1", 5000);
            System.out.println("Connected");

            DataInputStream in = new DataInputStream(s.getInputStream());

            DataOutputStream out = new DataOutputStream(s.getOutputStream());
            BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
            String m = ""; String m2 = "";

            while (!m.equals("Over")) {

                m = br.readLine();
                out.writeUTF(m);

                m2 = in.readUTF();
                System.out.println(m2);
                out.flush();

            }
            in.close();
            out.close();
            s.close();
        } catch (Exception i) {
            System.out.println(i);
        }

    }
}