package main

import (
    "flag"
    "fmt"
    "net"
    "os"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    // Parse comand line option for flags
    wordPtr := flag.String("hostname","localhost","Hostname to Listne on")
    portPtr := flag.String("port","3333","Port number to listen on")
    protPtr := flag.String("protocol","tcp","Proto - defaults to tcp")
    flag.Parse()
    fmt.Printf("Will listen on %v:%v/%v\n", *wordPtr,*portPtr,*protPtr)


    // Listen for incoming connections.
    l, err := net.Listen(*protPtr, *wordPtr +":"+ *portPtr)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + *wordPtr + ":" + *portPtr)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  fmt.Printf("Mangling Connection from %v\n",conn.RemoteAddr().String())
  // Make a buffer to hold incoming data.
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  reqLen, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }
  junky := fmt.Sprintf("Message (%v)'%v' received.",reqLen,buf)

  // Send a response back to person contacting us.
  conn.Write([]byte(junky))
  // Close the connection when you're done with it.
  conn.Close()
}
