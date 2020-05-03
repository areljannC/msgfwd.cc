<script>
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()

  let name = ''
  let location = ''
  let message = ''

  const handleSubmit = async () => {
    await fetch('http://localhost:8000/api/message', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        name,
        location,
        message
      })
    })
      .then((res) => {
        name = location = message = ''
        dispatch('submit')
      })
      .catch((err) => console.log(err))
  }
</script>

<style scoped>
  form {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
  }
</style>

<div>
  <form on:submit|preventDefault={handleSubmit}>
    <input bind:value={name} type="text" />
    <input bind:value={location} type="text" />
    <textarea bind:value={message} />
    <button type="submit">Submit</button>
  </form>
</div>
