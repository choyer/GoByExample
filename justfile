# create scaffolding for new exercise
lesson name:
  @echo 'Creating exercise: {{name}}...'
  mkdir {{name}}
  touch {{name}}/{{name}}.go
  touch {{name}}/NOTES.md

# Go run the lesson
run lesson:
  go run {{lesson}}/{{lesson}}.go
