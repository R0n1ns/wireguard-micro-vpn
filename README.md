<h1>WireGuard Micro VPN - Description and Setup Guide</h1>

  <h2>Program Description</h2>
  <p>
      The program is a command-line tool to manage a WireGuard VPN server and its clients. It provides functionalities to start a WireGuard VPN, add and manage clients, and send configuration details to clients via Telegram.
  </p>

  <p>
      The program uses the <a href="https://github.com/R0n1ns/wireguard_go_ubuntu" target="_blank">WireGuard Go Ubuntu</a> package for VPN configuration management and traffic collection.
  </p>

  <h2>Features</h2>
  <ul>
      <li>Start WireGuard with automatic or manual settings</li>
      <li>Add new clients to the VPN</li>
      <li>Activate, stop, and delete clients</li>
      <li>Send configuration data to clients via Telegram</li>
      <li>View all connected clients</li>
      <li>Remove and stop WireGuard configuration</li>
  </ul>

  <h2>Requirements</h2>
  <ul>
      <li>Go installed on your machine</li>
      <li>WireGuard installed on your server</li>
      <li>Telegram Bot token (for sending configuration data to clients)</li>
      <li>Clone the repository from <a href="https://github.com/R0n1ns/wireguard-micro-vpn" target="_blank">GitHub</a></li>
  </ul>

  <h2>How to Run the Program</h2>
  <ol>
      <li>Clone the repository:
          <pre><code>git clone https://github.com/R0n1ns/wireguard-micro-vpn.git</code></pre>
      </li>
      <li>Navigate to the project directory:
          <pre><code>cd wireguard-micro-vpn</code></pre>
      </li>
      <li>Build the project using Go:
          <pre><code>go build</code></pre>
      </li>
      <li>Run the program:
          <pre><code>./wireguard-micro-vpn</code></pre>
      </li>
      <li>
          Upon starting, the program will prompt you with several commands to manage the VPN:
          <ul>
              <li><strong>1</strong> - Start WireGuard (automatic or manual mode)</li>
              <li><strong>2</strong> - Add a new client</li>
              <li><strong>3</strong> - Stop a client</li>
              <li><strong>4</strong> - Activate a client</li>
              <li><strong>5</strong> - Delete a client</li>
              <li><strong>6</strong> - View all clients</li>
              <li><strong>7</strong> - Send client configuration data via Telegram</li>
              <li><strong>-2</strong> - Remove and stop WireGuard configuration</li>
              <li><strong>0</strong> - Show the command menu</li>
              <li><strong>-1</strong> - Exit the program</li>
          </ul>
      </li>
  </ol>

  <h2>Example Usage</h2>
  <p>After launching the program, you can follow these steps to start the VPN and add a new client:</p>
  <ol>
      <li>Enter <strong>1</strong> to start WireGuard.</li>
      <li>Select <strong>0</strong> for automatic setup or <strong>1</strong> for manual setup.</li>
      <li>If manual, enter the requested details such as the port number, server IP, and interface name.</li>
      <li>To add a new client, enter <strong>2</strong> and follow the prompts.</li>
  </ol>

  <h2>GitHub Repository</h2>
  <p>You can find the full source code at the following link: 
      <a href="https://github.com/R0n1ns/wireguard-micro-vpn" target="_blank">https://github.com/R0n1ns/wireguard-micro-vpn</a>
  </p>
