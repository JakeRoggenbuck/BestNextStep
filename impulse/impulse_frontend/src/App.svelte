<script>
  import { onMount } from "svelte";
  import Step from "./Step.svelte";
  import Select from "./Select.svelte";
  import "./app.css";

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

	  let cols;
	  let selected = 1;

  onMount(async () => {
    await fetch(`http://127.0.0.1:1357/api/v1/col/`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Basic " + btoa("Admin:banana"),
      },
    })
      .then((r) => r.json())
      .then((data) => {
        cols = JSON.parse(data["message"]);
      });
  });

</script>

<main>
	<h1>Best Next Step {selected}</h1>
<div>
  <select class="glass" name="membership" id="membership" bind:value={selected}>
    {#if cols}
      {#each cols as col}
        <option value={col._id}>{col.name}</option>
      {/each}
    {:else}
      <p class="loading">loading...</p>
    {/if}
  </select>
  <slot {selected}/>
</div>

  <div align="center">
    {#if steps}
      {#each steps as step}
        <Step {step} />
      {/each}
    {:else}
      <p class="loading">loading...</p>
    {/if}
  </div>
</main>

<style>
  :global(body) {
    background: rgb(174, 228, 238);
    background: linear-gradient(
      90deg,
      rgba(174, 228, 238, 1) 0%,
      rgba(148, 187, 233, 1) 100%
    );
  }

  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;

    background: rgba(255, 255, 255, 0.25);
    box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
    backdrop-filter: blur(4px);
    -webkit-backdrop-filter: blur(4px);
    border-radius: 10px;
    border: 1px solid rgba(255, 255, 255, 0.18);
    margin: 4px;
  }

  h1 {
    margin-top: 0px;
    color: #0b0d21;
    font-size: 4em;
    font-weight: 700;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
