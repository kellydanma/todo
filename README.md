# todo

Small, fast CLI `todo` list.

## Getting started

Install Go (1.14+) and `go get` this project:

```
go get github.com/kellydanma/todo
go install github.com/kellydanma/todo/cmd/todo
```

You should have the `todo` binary in your `$GOPATH`, you're good to _go_!

## Usage

### Adding tasks

Let's add a few tasks:

```
$ todo -a "social distance"
$ todo -a "make some food"
$ todo -a "eat my food!"
```

### Listing tasks

Let's see what our list looks like:

```
$ todo -l
 1. social distance                [ ]
 2. make some food                 [ ]
 3. eat my food!                   [ ]
```

#### Enable verbose output

For more detailed output, use the `-v` flag:

```
$ todo -l -v
 1. social distance                [ ] created on 2020-05-13, not complete
 2. make some food                 [ ] created on 2020-05-13, not complete
 3. eat my food!                   [ ] created on 2020-05-13, not complete
```

### Completing tasks

To mark a task as complete, use the `-c` flag and enter the task number:

```
$ todo -c 2
$ todo -l -v
 1. social distance                [ ] created on 2020-05-13, not complete
 2. make some food                 [✓] created on 2020-05-13, completed on 2020-05-13
 3. eat my food!                   [ ] created on 2020-05-13, not complete
```

#### Hide completed tasks

If you don't want to see completed tasks, use the `-h` flag:

```
$ todo -l -h -v
 1. social distance                [ ] created on 2020-05-13, not complete
 3. eat my food!                   [ ] created on 2020-05-13, not complete
```

### Deleting tasks

```
$ todo -l
 1. social distance                [ ]
 2. make some food                 [✓]
 3. eat my food!                   [ ]
$ todo -d 1
$ todo -l -v
 1. make some food                 [✓] created on 2020-05-13, completed on 2020-05-13
 2. eat my food!                   [ ] created on 2020-05-13, not complete
```
