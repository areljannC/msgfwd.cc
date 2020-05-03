<script>
  import { onMount } from 'svelte'
  import Form from '../components/Form.svelte'

  let isLoading = false
  let latestMessage = {}

  const getLatestMessage = async () => {
    if (process.browser) {
      isLoading = true
      await fetch('http://localhost:8000/api/message')
        .then(res => res.json())
        .then(data => {
          latestMessage = data
          isLoading = false
        })
        .catch(err => {
          console.log(err)
          isLoading = false
        })
    }
  }

  onMount(async () => {
    getLatestMessage()
  })
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
  {#if isLoading}
    <p>Getting latest message...</p>
  {:else}
    <p>Name: {latestMessage.name}</p>
    <p>Location: {latestMessage.location}</p>
    <p>Message: {latestMessage.message}</p>
  {/if}

  <Form on:submit={() => getLatestMessage()} />
</div>
