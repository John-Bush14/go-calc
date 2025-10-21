mod macros;
mod infix;
use derive_more::Display;


pub trait Tokenizer {
    fn tokenize(str: &str) -> Option<Vec<Token>>;

    fn is_valid_equation(str: &str) -> bool {Self::tokenize(str).is_some()}
}


#[derive(Display)]
#[display("{_0}")]
pub enum Token {
    Constant(f64),
    Operator(Operator),
    Bracket(Side)
}

#[derive(Display)]
pub enum Side {
    #[display("(")]
    Left,
    #[display(")")]
    Right
}

define_operators!(x, y, pub Operator {
    Addition("+") = x+y,
    Negation("-") = x-y,
    Multiplication("*") = x*y,
    Division("/") = x/y,
    Exponentiation("^") = x.powf(y),
    RootExtraction("√") = x.powf(1f64/y)
});


#[test]
fn token_formating_tests() {
    assert_eq!(format!("{}", Token::Constant(13f64)), "13");
    assert_eq!(format!("{}", Token::Operator(Operator::Multiplication)), "*");
    assert_eq!(format!("{}", Token::Bracket(Side::Left)), "(");
}
