package buffer

type Mode string

const (
    NORMAL Mode = "NORMAL"
    INSERT Mode = "INSERT"
)

type Buffer struct {
    TermCol int
    TermLine int
    Width int
    Height int
    Lines []Line
    CurrentLine *Line
    StatusLine StatusLine
    fillLine Line
    CursorLine int
    CursorCol int
    heldCursorCol int
}

func NewBuffer(termCol int, termLine int, width int, height int) Buffer {

    lines := []Line{BlankLine()}

    currLine := &lines[0]
    statusLine := StatusLine{"", "", " | ", NORMAL, 0, 0, "", 0, " "}
    fillLine := FillLine(width)

    buffer := Buffer{
        termCol,
        termLine,
        width,
        height,
        lines,
        currLine,
        statusLine,
        fillLine,
        0,
        0,
        0,
    }

    return buffer
}

