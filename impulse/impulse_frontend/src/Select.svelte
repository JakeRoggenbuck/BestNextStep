<script>
  import "./app.css";
  import { onMount } from "svelte";

	  let cols;
	  let selected;

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

<div>
  <select class="glass" name="membership" id="membership" bind:value={selected} on:change="{() => alert(selected)}">
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
