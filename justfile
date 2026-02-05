# create scaffolding for new exercise/lesson
new lesson:
  @echo 'Creating exercise: {{lesson}}...'
  mkdir {{lesson}}
  touch {{lesson}}/{{lesson}}.go
  touch {{lesson}}/NOTES.md
  echo "#{{lesson}} NOTES" > {{lesson}}/NOTES.md

# Go run the lesson
run lesson:
  go run {{lesson}}/{{lesson}}.go
