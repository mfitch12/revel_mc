# Revel Mail Congress
Messing around with Revel and Go

# Install dependencies

This project uses Revel framework. You can install it by:

<code>go get github.com/revel/revel</code>

<code>go get github.com/revel/cmd/revel</code>

# Run

Once Revel framework is installed, you can run the server by:

<code>revel run mailCongress</code>

Note that the project must be located under <code>$GOPATH/src/mailCongress</code>

# Routes

The API routes are defined in <code>conf/routes</code> file:

<code>/</code> Displays default hello world app

<code>GET /representatives/{zipcode}</code> Retrieve representatives by zip code {zipcode} (just returns dummy data so far)
