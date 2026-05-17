import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;
import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;

class HelloHandler implements HttpHandler {

    @Override
    public void handle(HttpExchange exchange) throws IOException {
        exchange.getResponseHeaders().set("Content-Type", "text/plain");
        exchange.sendResponseHeaders(200, 0);
        OutputStream od = exchange.getResponseBody();
        String response = "Hello, I'm Rafiz.";
        od.write(response.getBytes());
        od.close();
        exchange.close();
        
    }
}

public class SimpleHttpServer {
    public static void main(String[] args) throws IOException {
        HttpServer server = HttpServer.create(new InetSocketAddress(8080), 0);
        server.createContext("/rafiz", new HelloHandler());

        server.start();
        System.out.println("Server started and listening on port 8080");
    }
}