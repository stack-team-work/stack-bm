import { ref } from 'vue'
import { useMessage } from 'naive-ui'

export function useModal() {
  const message = useMessage()
  const showModal = ref(false)
  const isEdit = ref(false)
  const editId = ref(null)
  const submitLoading = ref(false)
  const formRef = ref(null)

  function open() {
    isEdit.value = false
    editId.value = null
    showModal.value = true
  }

  function openEdit(row) {
    isEdit.value = true
    editId.value = row.id
    showModal.value = true
  }

  function close() {
    showModal.value = false
  }

  async function submit(formData, createFn, updateFn) {
    try {
      await formRef.value?.validate()
    } catch {
      return false
    }

    submitLoading.value = true
    try {
      if (isEdit.value) {
        await updateFn(editId.value, { ...formData })
        message.success('更新成功')
      } else {
        await createFn({ ...formData })
        message.success('创建成功')
      }
      showModal.value = false
      return true
    } catch (err) {
      message.error(err.message || '操作失败')
      return false
    } finally {
      submitLoading.value = false
    }
  }

  async function handleDelete(id, deleteFn) {
    try {
      await deleteFn(id)
      message.success('删除成功')
      return true
    } catch (err) {
      message.error(err.message || '删除失败')
      return false
    }
  }

  return { showModal, isEdit, editId, submitLoading, formRef, open, openEdit, close, submit, handleDelete }
}
