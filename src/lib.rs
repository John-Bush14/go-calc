use derive_more::Display;
use thiserror::Error;
mod macros;
mod infix;


pub trait Tokenizer {
    fn tokenize(str: &str) -> Result<Vec<Token>, TokenizationError>;

    fn is_valid_equation(str: &str) -> Option<TokenizationError> {Self::tokenize(str).err()}
}

#[derive(Error, Debug)]
pub enum TokenizationError {
    #[error("Invalid token found in equation: '{0}'")]
    InvalidToken(String),
    #[error("Token '{0}' at index {1} has been found to be in surplus.")]
    Surplus(Token, usize),
}


#[derive(Display, Debug)]
#[display("{_0}")]
pub enum Token {
    Constant(f64),
    Operator(Operator),
    Bracket(Side)
}

#[derive(Display, Debug)]
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
