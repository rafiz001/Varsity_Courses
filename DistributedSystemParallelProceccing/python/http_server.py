from http.server import HTTPServer, BaseHTTPRequestHandler
import os

class PDFHandler(BaseHTTPRequestHandler):
    
    def do_GET(self):
        if self.path == '/pdf':
            self.handle_pdf()
        elif self.path == '/html':
            self.handle_html()
        else:
            self.send_response(404)
            self.end_headers()
            self.wfile.write(b'404 Not Found')
    
    def handle_pdf(self):
        try:
            # Update this path to your PDF file location
            pdf_path = "/home/rafiz/Downloads/tmp/report-1778721537160.pdf"
            # Set response headers
            self.send_response(200)
            self.send_header('Content-Type', 'application/pdf')
            self.end_headers()
            
            # Read and send the PDF file in chunks
            with open(pdf_path, 'rb') as file:
                self.wfile.write(file.read())
                
        except Exception as e:
            self.send_response(500)
            self.end_headers()
            self.wfile.write(f'Error: {str(e)}'.encode())
    
    def handle_html(self):
        html_response = """
            <html>
                <head><title>PDF Server</title></head>
                <body style="font-family: Arial; text-align:center; margin-top:50px;">
                    <h1>Welcome to My PDF supported HTTP Server</h1>
                    <p><u>Do you want today lab report cover page?.</u></p>
                    <a href="/pdf">Click here to view the cover page</a>
                </body>
            </html>
        """
        
        self.send_response(200)
        self.send_header('Content-Type', 'text/html')
        self.end_headers()
        self.wfile.write(html_response.encode())

port = 8080
server_address = ('', port)
httpd = HTTPServer(server_address, PDFHandler)
print(f'Server started on port {port}')
print(f'Visit: http://localhost:{port}/html')
httpd.serve_forever()

