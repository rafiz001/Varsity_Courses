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
        String response = "Hello, Rafiz.";
        od.write(response.getBytes());
        od.close();
        exchange.close();
        
    }
}
class Html implements HttpHandler {

    @Override
    public void handle(HttpExchange exchange) throws IOException {
        exchange.getResponseHeaders().set("Content-Type", "text/html");
        exchange.sendResponseHeaders(200, 0);
        OutputStream od = exchange.getResponseBody();
        String response = """
        <html>
        <body>rafiz <marquee>im rafiz</marquee>
        </body>
        </html>
        """;
        od.write(response.getBytes());
        od.close();
        exchange.close();
        
    }
}

public class SimpleHttpServer {

    public static void main(String[] args) throws IOException {
        HttpServer server = HttpServer.create(new InetSocketAddress(8080), 0);
        server.createContext("/rafiz", new HelloHandler());
        server.createContext("/", new Html());

        server.start();
        System.out.println("Server started and listening on port 8080");
    }
}