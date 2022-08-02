<script>
  import "./app.css";
  export let step;
  import {
    Trash2Icon,
    EditIcon,
    ArrowUpIcon,
    ArrowDownIcon,
    SquareIcon,
    CheckSquareIcon,
  } from "svelte-feather-icons";

  function toggleDone() {
    step.done = !step.done;
  }

  async function deleteStep(step_id) {
    await fetch(`http://127.0.0.1:1357/api/v1/step/` + step_id + "/", {
      method: "DELETE",
      headers: {
        // Test Passowrd for running locally - not used in production
        Authorization: "Basic " + btoa("Admin:banana"),
      },
    })
      .then((r) => r.json())
      .then((data) => {
        console.log(data);
      })
      .catch((error) => {
        console.log(error);
      });
  }
</script>

<div class="step flat">
  <div class="up" on:click={() => alert("up")}>
    <ArrowUpIcon size="24" />
  </div>

  <div class="down" on:click={() => alert("down")}>
    <ArrowDownIcon size="24" />
  </div>

  <div class="delete" on:click={async () => deleteStep(step._id)}>
    <Trash2Icon size="24" />
  </div>

  <div class="edit" on:click={() => alert("edit")}>
    <EditIcon size="24" />
  </div>

  <div class="done" on:click={() => toggleDone()}>
    {#if step.done}
      <CheckSquareIcon class="checked" size="24" />
    {:else}
      <SquareIcon class="unchecked" size="24" />
    {/if}
  </div>

  <h3>{step._id}. {step.name}</h3>
  <p class="desc">{step.desc}</p>
</div>
