mod macros;

use crate::define_operators;
use derive_more::Display;


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

mod lexer;
