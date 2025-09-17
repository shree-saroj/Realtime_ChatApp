var socket = new WebSocket("ws://192.168.100.4:8080/ws");

let connect = (cb) => {
  console.info("Attempting Connection...");

  socket.onopen = () => {
    console.info("Successfully Connected!!!");
  };
  socket.onmessage = (message) => {
    cb(message);
  };

  socket.onclose = () => {
    console.info("Connection Closed...");
  };

  socket.onerror = (error) => {
    console.error("WebSocket Error:", error);
  };
};

let sendMsg = (msg) => {
  console.info("Sending Message:" + msg);
  socket.send(msg);
};

export { connect, sendMsg };
