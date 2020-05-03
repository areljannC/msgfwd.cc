<script>
  const getData = async () =>
    await fetch('https://jsonplaceholder.typicode.com/users')
      .then(res => res.json())
      .then(data => data)
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
  {#await getData()}
    <p>fetching data...</p>
  {:then data}
    {#each data as user, i (user.id)}
      <p>User: {user.name}</p>
    {/each}
  {/await}
</div>
