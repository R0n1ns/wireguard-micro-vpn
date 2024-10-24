<h1>WireGuard Micro VPN</h1>
  <p><a href="https://github.com/R0n1ns/wireguard-micro-vpn">GitHub Repository</a></p>

  <h2>Compiled Program Description</h2>
  <p>
      The WireGuard Micro VPN is a simple command-line application written in Go that allows users to manage WireGuard VPN configurations. 
      Users can configure server settings, add clients, start/stop VPN connections, and send configuration files to clients via Telegram. 
      The program operates with both automatic and manual modes for setting up WireGuard server configurations.
  </p>

  <h3>Main Features:</h3>
  <ul>
      <li>Start WireGuard server in automatic or manual configuration mode.</li>
      <li>Generate and display server private and public keys.</li>
      <li>Add new clients to the VPN with unique keys and addresses.</li>
      <li>Send client configuration files via Telegram using bot integration.</li>
      <li>Start, stop, or delete client connections individually.</li>
      <li>View all configured clients and their connection details.</li>
      <li>Remove the WireGuard configuration and scripts, stopping the VPN server.</li>
  </ul>

  <h2>Instructions for Running the Compiled Program</h2>
  <ol>
      <li><strong>Clone the repository:</strong></li>
      <pre><code>git clone https://github.com/R0n1ns/wireguard-micro-vpn</code></pre>

      <li><strong>Install dependencies:</strong> Make sure you have Go installed on your system.</li>
      <pre><code>go get github.com/R0n1ns/wireguard_go_ubuntu</code></pre>

      <li><strong>Build the program:</strong></li>
      <pre><code>go build -o wireguard-micro-vpn main.go</code></pre>

      <li><strong>Run the compiled program:</strong></li>
      <pre><code>./wireguard-micro-vpn</code></pre>

      <li><strong>Program Usage:</strong> After starting the program, follow the prompts to interact with the WireGuard VPN setup. The available commands include:</li>
      <ul>
          <li>1: Start WireGuard server (manual or automatic configuration)</li>
          <li>2: Add a new client</li>
          <li>3: Stop a client connection</li>
          <li>4: Activate a client connection</li>
          <li>5: Delete a client connection</li>
          <li>6: List all clients</li>
          <li>7: Send client configuration to Telegram</li>
          <li>-2: Remove WireGuard configuration and stop the server</li>
          <li>-1: Exit the program</li>
      </ul>
  </ol>

  <h3>Automatic Configuration Mode:</h3>
  <p>
      When starting the WireGuard server, you can choose automatic mode by entering <code>0</code> when prompted. 
      In this mode, the server will use default settings such as port 51820 and network interface <code>eth0</code>.
  </p>

  <h3>Manual Configuration Mode:</h3>
  <p>
      If you choose manual configuration by entering <code>1</code>, you will be prompted to enter the server's port, 
      network interface name, and server IP. You will also be required to provide a bot token for Telegram integration.
  </p>

  <h3>Telegram Integration:</h3>
  <p>
      The program uses a Telegram bot to send client configuration files directly to users. You must provide a valid bot token 
      when prompted during the setup or add client processes.
  </p>

  <h2>Error Handling</h2>
  <p>
      If an error occurs during the execution of the program, it will automatically attempt to recover and reload the saved configuration from <code>data.json</code>.
  </p>

  <h2>Exiting the Program</h2>
  <p>
      To safely exit the program, enter <code>-1</code> at any time. This will save the current state and terminate the program.
  </p>
