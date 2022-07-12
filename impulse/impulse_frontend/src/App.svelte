<script>
  import { onMount } from "svelte";

  let steps;

  onMount(async () => {
    await fetch(`http://127.0.0.1:1357/api/v1/step/`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Basic " + btoa("Admin:banana"),
      },
    })
      .then((r) => r.json())
      .then((data) => {
        steps = JSON.parse(data["message"]);
      });
  });
</script>

<main>
  <h1>Best Next Step</h1>
  <div align="center">
    {#if steps}
      {#each steps as step}
        <div class="step">
          <h2>{step._id}. {step.name}</h2>
          <h4>{step.desc}</h4>
          <h6>from collection {step.collection}</h6>
        </div>
      {/each}
    {:else}
      <p class="loading">loading...</p>
    {/if}
  </div>
</main>

<style>
  .step {
    background-color: #f5eef8;
    width: 50%;
    margin: 8px;
    padding: 2px;
  }

  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #9e8eff;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 300;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
