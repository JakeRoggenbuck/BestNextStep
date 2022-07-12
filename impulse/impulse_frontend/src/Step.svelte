<script>
  import "./app.css";
  export let step;
</script>

<div class="step glass">
  <h2>{step._id}. {step.name}</h2>
  <h4>{step.desc}</h4>
  <h6>from collection {step.collection}</h6>
</div>
