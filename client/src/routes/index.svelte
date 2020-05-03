<script>
  const getLatestMessage = async () =>
    process.browser
      ? await fetch('http://localhost:8000/api/message')
          .then(res => res.json())
          .then(data => data)
          .catch(err => console.log(err))
      : false
</script>

<style scoped>
  .container {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
</style>

<svelte:head>
  <title>msgfwd.cc | Forward Your Message</title>
</svelte:head>

<div class="container">
  {#await getLatestMessage()}
    <p>Getting latest message...</p>
  {:then {name, location, message}}
    <p>Name: {name}</p>
    <p>Location: {location}</p>
    <p>Message: {message}</p>
  {:catch error}
    <p>Oops... something happened with our server.</p>
  {/await}
</div>
