<script>
  async function testBackend() {
    try {
      const response = await fetch("/api/hello");
      const data = await response.json();
      alert(data.message);
    } catch (error) {
      console.error("Error connecting to backend: ", error);
      alert("Could not connect to backend. Is it running?");
    }
  }

  testBackend();
</script>

<main>
  <div>
    <h1>Hello World!</h1>
  </div>
</main>
