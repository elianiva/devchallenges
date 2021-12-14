<script>
  import PictureImage from "./Picture.svelte";

  export let dropHandler, fileInput, upload;
</script>

<h1 class="upload__title">Upload your image</h1>
<span class="upload__subtitle">File should be Jpeg, Png...</span>
<div
  class="upload__upload"
  on:dragover|preventDefault={() => console.log("YO")}
  on:drop|preventDefault={dropHandler}
>
  <PictureImage />
  <span class="upload__upload-detail"> Drag & Drop your image here </span>
</div>
<span class="upload__or">or</span>
<div class="upload__button-wrapper">
  <input
    bind:this={fileInput}
    on:change={(e) => upload(e.currentTarget.files[0])}
    class="upload__button-input"
    type="file"
  />
  <button class="upload__button" on:click={() => fileInput.click()}>
    Choose a file
  </button>
</div>

<style>
  .upload__title {
    font-family: "Poppins", sans-serif;
    font-size: 1.25rem;
    color: #4f4f4f;
  }

  .upload__subtitle {
    display: block;
    font-family: "Poppins", sans-serif;
    font-size: 0.9rem;
    color: #828282;
    margin: 1rem 0 1.5rem;
    font-weight: 500;
  }

  .upload__upload {
    position: relative;
    background-color: #f6f8fb;
    border-radius: 0.5rem;
    padding: 3rem 5rem;
    border: 2px #97bef4 dashed;
  }

  .upload__upload-detail {
    display: block;
    font-family: "Poppins", sans-serif;
    font-size: 0.8rem;
    color: #bdbdbd;
    margin-top: 1.5rem;
  }

  .upload__or {
    display: block;
    font-family: "Poppins", sans-serif;
    color: #bdbdbd;
    margin: 1.25rem 0;
  }

  .upload__button {
    font-family: "Poppins", sans-serif;
    font-size: 0.9rem;
    font-weight: 500;
    background-color: #2f80ed;
    color: #ffffff;
    border-radius: 0.5rem;
    border: none;
    outline: none;
    padding: 0.5rem 1rem;
    cursor: pointer;
  }

  .upload__button:hover {
    filter: brightness(0.95);
  }

  .upload__button-input {
    display: none;
  }
</style>
