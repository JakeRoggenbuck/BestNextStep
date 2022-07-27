<script>
  import "./app.css";
	  export let step;
	  import { Trash2Icon, EditIcon, ArrowUpIcon, ArrowDownIcon } from 'svelte-feather-icons';
</script>

<div class="step flat">

  <div class="up" on:click={() => alert("up")}>
	<ArrowUpIcon size="24" />
  </div>

  <div class="down" on:click={() => alert("down")}>
	<ArrowDownIcon size="24" />
  </div>

  <div class="delete" on:click={() => alert("delete")}>
	<Trash2Icon size="24" />
  </div>

  <div class="edit" on:click={() => alert("edit")}>
	<EditIcon size="24" />
  </div>
  <h3>{step._id}. {step.name}</h3>
  <p class="desc">{step.desc}</p>
</div>
