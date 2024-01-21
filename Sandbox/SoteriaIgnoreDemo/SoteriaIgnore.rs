use std::fs::read_to_string;

fn remove_comments_and_empty_lines(lines: Vec<String>) -> Vec<String> {
    let mut result = Vec::new();

    for line in lines {
        let line_without_comment: String = line.splitn(2, '#').next().unwrap_or("").trim().to_string();

        if !line_without_comment.is_empty() {
            result.push(line_without_comment);
        }
    }

    result
}

fn tokenize_lines(lines: &Vec<String>) -> Vec<String> {
    let mut token_results = Vec::new();

    for line in lines {
        if line.starts_with("-") {
            if line.ends_with("/") {
                token_results.push(line.to_string() + " : SkipFolder");
            }
            token_results.push(line.to_string() + " : SkipFile");
        }

        if line.starts_with("+") {
            if line.ends_with("/") {
                token_results.push(line.to_string() + " : AddFolder");
            }
            token_results.push(line.to_string() + " : AddFile");
        }

        if line.starts_with("*") {
            token_results.push(line.to_string() + " : SkipExtension");
        }
    }

    token_results
}

fn read_lines(filename: &str) -> Vec<String> {
    let mut result = Vec::new();
    let mut skip_multiline_comment = false;
    for line in read_to_string(filename).unwrap().lines() {
       
        if let Some(first_char) = line.chars().next() {
            if first_char == '#' {
                continue;
            }

            if line.starts_with("|-start") {
                skip_multiline_comment = true;
                continue;
            }

            if skip_multiline_comment && line.starts_with("|-end") {
                skip_multiline_comment = false;
                continue;
            }

            if skip_multiline_comment {
                continue;
            }
        }

        result.push(line.to_string())
    }
    
    let next_result = remove_comments_and_empty_lines(result);

    let next_next_result = tokenize_lines(&next_result);

    next_next_result
}


fn main() {
    let ignore_file = "./.soteriaignore";
    println!("File Name: {}", ignore_file);

    let content = read_lines(ignore_file);

    println!("Content:\n{:?}", content);
}
