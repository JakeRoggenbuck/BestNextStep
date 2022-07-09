<script>
  import { onMount } from 'svelte';
  let steps;
  onMount(async () => {
    await fetch(`http://127.0.0.1:1357/`, {
      method: 'GET',
      headers: {
        Authorization: 'Basic ' + btoa('Admin:banana'),
      },
    }).then((data) => {
      steps = data;
    });
  });
</script>

<main>
  {#if steps}
    {steps}
  {:else}
    <p class="loading">loading... ahh</p>
  {/if}
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
