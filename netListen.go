package main

import (
    "flag"
    "fmt"
    "net"
    "os"
    "time"
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
    //FIXME - Validate Arguments ???


    returnCode := ListenLoop(*wordPtr,*portPtr,*protPtr,10)
    if (returnCode !=0) {
        //Try Try again?
        fmt.Println("Exiting normally on abnormal code ", returnCode)
        os.Exit(0)
    }
    fmt.Println("Exiting normally on normal code ", returnCode)

}


//Main Loop After Arg Parse
func ListenLoop(host,port,prot string, timeout int) int {
  if (prot != "tcp") {
    return -99
  }
  var junky error
  fmt.Printf("Will listen on %v:%v/%v\n", host,port,prot)
  resolvedAddr, junky := net.ResolveTCPAddr(prot,host +":"+ port)
  if (junky != nil){
    fmt.Println("Error resolving:", junky.Error())
    return -99
  }
  // Listen for incoming connections.
  l, junky := net.ListenTCP(prot, resolvedAddr)
  if junky != nil {
      fmt.Println("Error listening:", junky.Error())
      return -1
  }

  junky = l.SetDeadline(time.Now().Add(time.Duration(timeout)*time.Second))
  if junky != nil {
      fmt.Println("Error Setting Timeout ",timeout," - ", junky.Error())
      return -1
  }

  // Close the listener when the application closes.
  // FIXME - is this still valid when moved to a function?
  defer l.Close()
  fmt.Println("Actually Listening on " + host + ":" + port)
  for {
      // Listen for an incoming connection.
      conn, err := l.Accept()
      if err != nil {
        if err, ok := err.(*net.OpError); ok && err.Timeout() {
          // it was a timeout
          return 0
        }
        fmt.Println("Error accepting: ", err.Error())
        return -2

      }
      // Handle connections in a new goroutine.
      go handleRequest(conn)
  }
  return 0
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
