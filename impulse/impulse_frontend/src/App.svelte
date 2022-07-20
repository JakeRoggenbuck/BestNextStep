<script>
  import { onMount } from "svelte";
  import Step from "./Step.svelte";
  import "./app.css";

  let cols;
  let selected = 1;

  onMount(async () => {
    await fetch(`http://127.0.0.1:1357/api/v1/col/`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        // Test Passowrd for running locally - not used in production
        Authorization: "Basic " + btoa("Admin:banana"),
      },
    })
      .then((r) => r.json())
      .then((data) => {
        cols = JSON.parse(data["message"]);
      });
  });

  let steps;
  let current;

  onMount(async () => {
    await fetch(`http://127.0.0.1:1357/api/v1/step/`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        // Test Passowrd for running locally - not used in production
        Authorization: "Basic " + btoa("Admin:banana"),
      },
    })
      .then((r) => r.json())
      .then((data) => {
        steps = JSON.parse(data["message"]);

        current = steps.filter(inCollection);
      });
  });

  function inCollection(x) {
    return x["collection"] == selected;
  }

  function setCurrent() {
    current = steps.filter(inCollection);
  }
</script>

<main>
  <div align="center">
    <div class="glass contain">
      <h2>Best Next Step</h2>
      <div>
        <select
          class="flat"
          name="membership"
          id="membership"
          bind:value={selected}
          on:change={() => setCurrent()}
        >
          {#if cols}
            {#each cols as col}
              <option value={col._id}>{col.name}</option>
            {/each}
          {:else}
            <p class="loading">loading...</p>
          {/if}
        </select>
        <slot {selected} />
      </div>

      <div align="center">
        {#if current}
          {#each current as step}
            <Step {step} />
          {/each}
          <div class="glass step">
            <br />
            <form>
              <input
                class="flat title"
                type="text"
                id="title"
                name="title"
                placeholder="title..."
              />
              <input
                class="flat desc"
                type="text"
                id="desc"
                name="desc"
                placeholder="description..."
              />
              <input class="flat submit" type="submit" value="+" />
            </form>
            <br />
          </div>
        {:else}
          <p class="loading">loading...</p>
        {/if}
      </div>
    </div>
  </div>
</main>

<div class="footer">
  <a href="https://BestNextStep.org">BestNextStep.org</a> - Jake Roggenbuck -
  <a href="https://jr0.org">jr0.org</a>
  - <a href="https://github.com/jakeroggenbuck/BestNextStep">Source Code</a>
</div>

<style>
  :global(body) {
    background: rgb(174, 228, 238);
    background: linear-gradient(
      90deg,
      rgba(174, 228, 238, 1) 0%,
      rgba(148, 187, 233, 1) 100%
    );
  }

  .footer {
    font-size: 20px;
    margin: 10px;
  }

  main {
    text-align: center;
    padding: 1em;
    margin: 0 auto;
  }

  h2 {
    margin-top: 0px;
    margin-bottom: 5px;
    color: #0b0d21;
    font-size: 3em;
    font-weight: 700;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
