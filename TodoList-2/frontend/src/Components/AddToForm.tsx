import React, { FormEvent, useState } from 'react'
import {
  ActionGroup,
  Button,
  Form,
  FormGroup,
  TextContent,
  TextInput,
  TextVariants,
  Text,
  PageSectionVariants,
} from '@patternfly/react-core'

import {
  Page,
  Masthead,
  MastheadMain,
  MastheadBrand,
  PageSection,
} from '@patternfly/react-core'

interface Props {
  handleAdd: (todo: string) => void
}
const AddToForm = ({ handleAdd }: Props) => {
  const [todo, setTodo] = useState<string>('')

  const addTask = (e: FormEvent<HTMLButtonElement>) => {
    e.preventDefault()
    handleAdd(todo)
    setTodo('')
  }

  const handleChange = (task: string) => {
    setTodo(task)
  }

  const header = (
    <Masthead>
      <MastheadMain>
        <MastheadBrand href="https://patternfly.org" target="_blank">
          <TextContent
            style={{ textAlign: 'center', fontWeight: 'bold', padding: '2rem' }}
          >
            <Text component={TextVariants.h2}>TODO LIST</Text>
          </TextContent>
        </MastheadBrand>
      </MastheadMain>
    </Masthead>
  )

  return (
    <>
      <Page header={header}>
        <PageSection
          isWidthLimited
          isCenterAligned
          variant={PageSectionVariants.light}
        >
          <Form isWidthLimited method="post">
            <FormGroup label="Task" isRequired fieldId="0-label1">
              <TextInput
                isRequired
                id="0-label1"
                name="0-label1"
                value={todo}
                placeholder="Enter a task"
                onChange={handleChange}
              />
            </FormGroup>
            <ActionGroup>
              <Button variant="primary" type="submit" onClick={addTask}>
                Add
              </Button>
            </ActionGroup>
          </Form>
        </PageSection>
      </Page>
    </>
  )
}

export default AddToForm
