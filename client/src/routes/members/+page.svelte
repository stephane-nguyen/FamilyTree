<script lang="ts">
  import { onMount } from "svelte";
  import type { Member } from "$lib/types";
  import Spinner from "$lib/components/spinner.svelte";
  import Table from "./table.svelte";
  import Tree from "./tree.svelte";
  import AddMember from "./addMember.svelte";

  let members: Member[];
  let showTable: boolean = true;
  let showTree: boolean = false;
  let showAddMember: boolean = false;
  let error: { message: any };

  onMount(() => {
    getMembers()
      .then((data) => {
        members = data;
      })
      .catch((err) => {
        error = err;
      });
  });

  async function getMembers() {
    const response = await fetch("http://localhost:8000/v1/members");
    const json = await response.json();
    return json.members;
  }

  function handleTableVisibility() {
    showTable = true;
    showTree = false;
    showAddMember = false;
  }

  function handleTreeVisibility() {
    showTree = true;
    showTable = false;
    showAddMember = false;
  }

  function handleAddMemberVisibility() {
    showAddMember = true;
    showTree = false;
    showTable = false;
  }
</script>

<div class="px-3 py-2 bg-zinc-200">
  <button on:click={handleTableVisibility} type="button" class="p-2">
    <div class="rounded-lg hover:bg-zinc-300">
      <img src="/table-list-svgrepo-com.svg" alt="table" class="h-10 w-10" />
    </div>
  </button>

  <button on:click={handleTreeVisibility} type="button" class="p-2">
    <div class="rounded-lg hover:bg-zinc-300">
      <img src="/family-tree-svgrepo-com.svg" alt="family tree" class="h-10 w-10" />
    </div>
  </button>

  <button on:click={handleAddMemberVisibility} type="button" class="p-2">
    <div class="rounded-lg hover:bg-zinc-300">
      <img src="/add-profile-svgrepo-com.svg" alt="family tree" class="h-10 w-10" />
    </div>
  </button>
</div>

{#if showTable}
  {#if members}
    <Table {members} />
  {:else if error}
    <p style="color: red">{error.message}</p>
  {:else}
    <Spinner />
  {/if}
{/if}

{#if showTree}
  {#if members}
    <Tree />
  {:else if error}
    <p style="color: red">{error.message}</p>
  {:else}
    <Spinner />
  {/if}
{/if}

{#if showAddMember}
  <AddMember />
{/if}
