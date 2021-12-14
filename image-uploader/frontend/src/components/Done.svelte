<script>
  import Checkmark from "./Checkmark.svelte";
  import { isUploading, isFinished } from "../store";

  let linkElement;

  export let src;
</script>

<Checkmark className="done__mark" />
<h1 class="done__title">Uploaded Successfully!</h1>
<img class="done__image" alt={src} src="http://localhost:8000/img/{src}" />
<div class="done__link">
  <input
    class="done__link-text"
    bind:this={linkElement}
    type="text"
    readonly
    value="http://localhost:8000/img/{src}"
  />
  <button
    class="done__link-copy"
    on:click={() => {
      console.log(linkElement);
      linkElement.select();
      console.log(linkElement.innerText);
      document.execCommand("copy");
    }}>Copy Link</button
  >
</div>
<span
  class="done__more"
  on:click={() => {
    isUploading.set(false);
    isFinished.set(false);
  }}>Upload More</span
>

<style>
  :global(.done__mark) {
    width: 2.5rem;
    height: 2.5rem;
    color: #219653;
    margin-bottom: 0.5rem;
  }

  .done__title {
    font-family: "Poppins", sans-serif;
    font-size: 1.25rem;
    font-weight: 500;
    color: #4f4f4f;
    margin-bottom: 1rem;
  }

  .done__image {
    border-radius: 0.5rem;
    max-width: 26rem;
  }

  .done__link {
    background-color: #f6f8fb;
    border: 1px #e0e0e0 solid;
    padding: 0.25rem 0.25rem 0.25rem 0.5rem;
    margin-top: 0.5rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-radius: 0.5rem;
    min-width: 24rem;
  }

  .done__link-text {
    flex: 1;
    font-family: "Poppins", sans-serif;
    font-size: 0.75rem;
    color: #4f4f4f;
    font-weight: 500;
    background-color: transparent;
    border: none;
    outline: none;
  }

  .done__link-copy {
    background-color: #2f80ed;
    color: #ffffff;
    border-radius: 0.5rem;
    outline: none;
    border: none;
    padding: 0.5rem 0.75rem;
    cursor: pointer;
  }

  .done__link-copy:hover {
    filter: brightness(0.8);
  }

  .done__more {
    display: block;
    margin-top: 1rem;
    font-family: "Poppins", sans-serif;
    font-weight: 500;
    font-size: 0.8rem;
    color: #4f4f4f;
    cursor: pointer;
  }

  .done__more:hover {
    text-decoration: underline;
    color: #2f80ed;
  }
</style>
